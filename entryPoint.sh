#!/bin/sh
set -e
set -x

export password=$(cat /mnt/secrets-store/secrets | jq -r '.data.data.password')
echo $password

echo Namespace: $NAMESPACE
echo ElbName: $ELB_NAME

DYNAMIC_ELB_EXTERNAP_IP=$(kubectl get services $ELB_NAME -n $NAMESPACE -o json | jq -r ".status.loadBalancer.ingress[0].ip")
until [ ! -z $DYNAMIC_ELB_EXTERNAP_IP  ] && [ $DYNAMIC_ELB_EXTERNAP_IP != null ]
do
  echo DYNAMIC_ELB_EXTERNAP_IP: $DYNAMIC_ELB_EXTERNAP_IP
  sleep 30
  DYNAMIC_ELB_EXTERNAP_IP=$(kubectl get services $ELB_NAME -n $NAMESPACE -o json | jq -r ".status.loadBalancer.ingress[0].ip")
done

echo $DYNAMIC_ELB_EXTERNAP_IP
export ELB_EXTERNAP_IP=$DYNAMIC_ELB_EXTERNAP_IP

# Run the web service on container startup.
/app/server