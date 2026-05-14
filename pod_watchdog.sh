#!/usr/bin/env bash
# pod_watchdog.sh — failure detector for the molequla ecology run on a RunPod.
# Runs on the pod, emits one stdout line per concerning event so neo can grep-tail
# it through SSH + Monitor for per-event chat notifications.
#
# Each event line has the format:
#   <UTC time> [watchdog] <element> <CLASS>: <detail>
#
# Classes:
#   FAIL              — crash signature in train.log (panic / Traceback / Killed / OOM / assert / NaN loss)
#   HEARTBEAT_STALE   — train.log mtime older than 5 min (organism stuck or dead)
#   RSS_HIGH          — process RSS exceeds 8 GB (per-organism threshold)
#   DEAD              — pid file points at a process that no longer exists
#   DISK_LOW          — pod disk free < 5 GB
#
# Usage:
#   bash pod_watchdog.sh [eco_dir]      # default /workspace/runs/eco
#
# Stops only on TERM/INT — let the caller decide when to kill it.

set -u
ECO_DIR="${1:-/workspace/runs/eco}"
ORGANISMS=(earth air water fire)
HB_WARN_SEC=300
RSS_HIGH_KB=8000000
DISK_LOW_GB=5
POLL_SEC=30

trap 'exit 0' INT TERM

stamp() { date -u +%H:%M:%SZ; }
emit()  { printf '%s [watchdog] %s\n' "$(stamp)" "$*"; }

# Track last seen FAIL line per organism so we don't re-emit the same one.
declare -A LAST_FAIL=()

while :; do
  NOW=$(date +%s)

  # Per-organism scans.
  for e in "${ORGANISMS[@]}"; do
    WORK="$ECO_DIR/work_$e"
    LOG="$WORK/train.log"
    PIDF="$WORK/org.pid"

    if [ ! -f "$LOG" ]; then
      emit "$e: no train.log at $LOG yet"
      continue
    fi

    # FAIL — scan recent tail for crash signatures, dedupe by content hash.
    NEW_FAIL=$(tail -c 32768 "$LOG" 2>/dev/null \
      | grep -E 'panic:|Traceback|Killed|SIGKILL|out of memory|OOM|runtime error|^assert|loss=nan|loss=NaN|loss=inf|fatal error|segmentation fault' \
      | tail -3 || true)
    if [ -n "$NEW_FAIL" ]; then
      HASH=$(printf '%s' "$NEW_FAIL" | md5sum 2>/dev/null | awk '{print $1}')
      if [ -z "${LAST_FAIL[$e]:-}" ] || [ "${LAST_FAIL[$e]}" != "$HASH" ]; then
        LAST_FAIL[$e]=$HASH
        # Emit one event line per matched signature line.
        while IFS= read -r line; do
          [ -n "$line" ] && emit "$e FAIL: $line"
        done <<< "$NEW_FAIL"
      fi
    fi

    # HEARTBEAT — log mtime older than HB_WARN_SEC.
    LOG_MTIME=$(stat -c %Y "$LOG" 2>/dev/null || stat -f %m "$LOG" 2>/dev/null || echo 0)
    AGE=$((NOW - LOG_MTIME))
    if [ "$AGE" -gt "$HB_WARN_SEC" ]; then
      emit "$e HEARTBEAT_STALE: log idle ${AGE}s (mtime $(date -u -d @$LOG_MTIME +%T 2>/dev/null || date -u -r $LOG_MTIME +%T 2>/dev/null))"
    fi

    # PID + RSS + DEAD.
    if [ -f "$PIDF" ]; then
      PID=$(cat "$PIDF" 2>/dev/null || echo "")
      if [ -n "$PID" ]; then
        if [ ! -d "/proc/$PID" ]; then
          emit "$e DEAD: pid $PID gone"
        else
          RSS_KB=$(awk '/^VmRSS:/{print $2}' "/proc/$PID/status" 2>/dev/null || echo "")
          if [ -n "$RSS_KB" ] && [ "$RSS_KB" -gt "$RSS_HIGH_KB" ]; then
            emit "$e RSS_HIGH: ${RSS_KB} kB (> ${RSS_HIGH_KB})"
          fi
        fi
      fi
    fi
  done

  # Pod-wide disk free check.
  FREE_GB=$(df -BG "$ECO_DIR" 2>/dev/null | awk 'NR==2{gsub("G","",$4); print $4}' || echo "")
  if [ -n "$FREE_GB" ] && [ "$FREE_GB" -lt "$DISK_LOW_GB" ]; then
    emit "pod DISK_LOW: ${FREE_GB}G free at $ECO_DIR (< ${DISK_LOW_GB}G)"
  fi

  sleep "$POLL_SEC"
done
