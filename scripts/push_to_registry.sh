#!/bin/bash
# SPDX-License-Identifier: Apache-2.0
# Copyright (c) 2020 Intel Corporation

REGISTRY=${EMCODOCKERREPO}
IMAGE=$1
TAG=${TAG}

push_to_registry() {
    M_IMAGE=$1
    M_TAG=$2
    echo "Pushing ${M_IMAGE} to ${REGISTRY}${M_IMAGE}:${M_TAG}..."
    docker tag ${M_IMAGE}:latest ${REGISTRY}${M_IMAGE}:${M_TAG}
    docker tag ${M_IMAGE}:latest ${REGISTRY}${M_IMAGE}:latest
    docker push ${REGISTRY}${M_IMAGE}:${M_TAG}
    docker push ${REGISTRY}${M_IMAGE}:latest
}

if [ "${BUILD_CAUSE}" != "TIMERTRIGGER" ] && [ "${BUILD_CAUSE}" != "DEV_TEST" ] && [ "${BUILD_CAUSE}" != "RELEASE" ]; then
    TAG="latest"
    exit 0
fi

if [ -z ${TAG} ]; then
  TAG=${BRANCH}-daily-`date +"%Y%m%d"`
fi

if [ "${BUILD_CAUSE}" == "DEV_TEST" ]; then
  TAG=${USER}-latest
fi

if [ "${BUILD_CAUSE}" == "TIMERTRIGGER" ] ; then
  if [ -z "${CI_COMMIT_REF_NAME}" ]; then
    CI_COMMIT_REF_NAME=${BRANCH}
  fi
  TAG=${CI_COMMIT_REF_NAME}-daily-`date +"%Y%m%d"`-${COMMITID}
fi

if [ "${BUILD_CAUSE}" == "RELEASE" ]; then
  if [ ! -z ${EMCOSRV_RELEASE_TAG} ]; then
    # remove first character of release tag, i.e. the "v"
    TAG=${EMCOSRV_RELEASE_TAG:1}
  else
    TAG=${TAG}
  fi
  if [ -z ${TAG} ]; then
    echo "HEAD has no tag associated with it"
    exit 0
  fi
fi

[[ -z "$MODS" ]] && export MODS="clm dcm dtc nps gac monitor ncm orch ovn rsync"
MODS=$(echo $MODS | sed 's;tools/emcoctl;;')

for m in ${MODS}; do
  push_to_registry emco-$m ${TAG}
done
