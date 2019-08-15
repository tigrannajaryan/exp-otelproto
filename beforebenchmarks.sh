#!/usr/bin/env bash

# Set performance governor for all CPUs
cpufreq-set -g performance -c 0
cpufreq-set -g performance -c 1
cpufreq-set -g performance -c 2
cpufreq-set -g performance -c 3
cpufreq-set -g performance -c 4
cpufreq-set -g performance -c 5
cpufreq-set -g performance -c 6
cpufreq-set -g performance -c 7

# Disable turboboost (see https://easyperf.net/blog/2019/08/02/Perf-measurement-environment-on-Linux)
echo 1 > /sys/devices/system/cpu/intel_pstate/no_turbo
