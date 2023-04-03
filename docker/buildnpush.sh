

# read tag from first argument or from git tag
if [ ! -z $1 ]
then
  TAG=$1
else
  TAG=`git describe --always`
fi

REGISTRY="registry.gitlab.com/khakibee/khakibee"

# build images from dockerfiles
cd ../api
docker build -t $REGISTRY/api-admin:$TAG -f Dockerfile.admin . &
docker build -t $REGISTRY/api-calendar:$TAG -f Dockerfile.calendar . &

cd ../ui-admin
docker build -t $REGISTRY/ui-admin:$TAG . &

cd ../ui-calendar
docker build -t $REGISTRY/ui-calendar:$TAG  . &

cd ../nginx
docker build -t $REGISTRY/nginx:$TAG . &


# wait for build to finish
wait


# push images to registry
docker push $REGISTRY/api-admin:$TAG &
docker push $REGISTRY/ui-admin:$TAG &
docker push $REGISTRY/ui-calendar:$TAG &
docker push $REGISTRY/api-calendar:$TAG &
docker push $REGISTRY/nginx:$TAG &


# wait for push to finish
wait

# remove images from local registry
docker rmi $REGISTRY/ui-admin:$TAG
docker rmi $REGISTRY/api-admin:$TAG
docker rmi $REGISTRY/ui-calendar:$TAG
docker rmi $REGISTRY/api-calendar:$TAG
docker rmi $REGISTRY/nginx:$TAG