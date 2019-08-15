#!/usr/bin/env bash

# Restore powersave governor
cpufreq-set -g powersave -c 0
cpufreq-set -g powersave -c 1
cpufreq-set -g powersave -c 2
cpufreq-set -g powersave -c 3
cpufreq-set -g powersave -c 4
cpufreq-set -g powersave -c 5
cpufreq-set -g powersave -c 6
cpufreq-set -g powersave -c 0

# Restore turboboost
echo 0 > /sys/devices/system/cpu/intel_pstate/no_turbo
