#!/usr/bin/env bash


errcho() {
    (>&2 echo -e "\e[31m$1\e[0m")
}

FAILED=

if [[ -n $(python -c "from vprint import vprint; vprint('nope')" 2>&1) ]]; then
  errcho "1. Saw output when nothing was expected"
  FAILED=1
fi


if [[ -n $(VPRINTV=1 python -c "from vprint import vprint; vprint('stderr_only')" 2>/dev/null) ]]; then
  errcho "2. Saw output on stdout when none was expected"
  FAILED=1
fi


if [[ -z $(VPRINTV=1 python -c "from vprint import vprint; vprint('yup')" 2>&1) ]]; then
  errcho "3. Saw no output when some was expected"
  FAILED=1
fi

if [[ -z $(VPRINTV=1 VPRINTV_FILE=stdout python -c "from vprint import vprint; vprint('yup')" 2>&1) ]]; then
  errcho "4. Saw no output when some was expected"
  FAILED=1
fi


if [[ -z $(VPRINTV=1 python -c "from vprint import vprint; vprint('yup')" 2>&1 1>/dev/null) ]]; then
  errcho "5. Expected to see stderr, did not see anything"
  FAILED=1
fi


if [[ -n $(VPRINTV=1 python -c "from vprint import vprint; vprint('yup')" 2>/dev/null ) ]]; then
  errcho "6. Expected to see stderr, did not see anything"
  FAILED=1
fi


if [[ -n "$FAILED" ]]; then
  echo "Failed one or more tests"
  exit 1
fi

echo "PASS!"