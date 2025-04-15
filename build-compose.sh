#!/bin/bash

# Parse command-line arguments
exclude_services=()
services_to_build=()
all_services=()

while [[ $# -gt 0 ]]; do
  case "$1" in
    --exclude)
      shift
      exclude_services+=("$1")
      shift
      ;;
    *)
      services_to_build+=("$1")
      shift
      ;;
  esac
done

# Get all services from docker-compose.yml
all_services=($(docker-compose config --services))

# If no services were specified, build all except excluded
if [[ ${#services_to_build[@]} -eq 0 ]]; then
  services_to_build=("${all_services[@]}")
fi

# Remove excluded services from the list
final_services=()
for service in "${services_to_build[@]}"; do
  if [[ ! " ${exclude_services[@]} " =~ " ${service} " ]]; then
    final_services+=("$service")
  fi
done

# Run docker-compose build with the final list of services
if [[ ${#final_services[@]} -gt 0 ]]; then
  docker-compose build "${final_services[@]}"
else
  echo "No services to build."
fi