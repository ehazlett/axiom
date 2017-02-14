# Axiom
Metadata service for Docker

# Usage
It is strongly recommended to use a separate network for the metadata service.
This allows one to attach containers to this network to access the service.

First, ensure that Swarm mode is enabled.  If not, run:

`$> docker swarm init`

Create a network:

`$> docker network create -d overlay axiom`

Create the Axiom service:

```
$> docker service create \
    --name axiom \
    --network axiom \
    --mount type=bind,src=/var/run/docker.sock,target=/var/run/docker.sock,ro=1 \
    ehazlett/axiom run
```

To test, create an example shell service:

```
$> docker service create -t \
    --name shell \
    --network axiom \
    --env LOOP=1 \
    ehazlett/curl axiom/containers
```

View the logs from the shell service (brace yourself for this horrendous one-liner):

```
$> docker logs --tail=1 \
    $(docker ps -q -f "label=com.docker.swarm.task.id=$(docker service ps -f "name=shell.1" -q shell)")
```

You should see JSON output of the containers:

```
[{"Id":"1f88e56da5dd2007e9b63f144f9f8810104e2be03e985a0224384ac12a46f72b","Names":["/shell.1.o6zf332yjf3idld9y3go7aell"],"Image":"ehazlett/curl@sha256:40e0cd8d8ee09d9ce445dc9d1ba29af32114b5c62daec394de93f9a5320cc622","ImageID":"sha256:c8127af118e02c6fe617a63490455073529f28ec0a6b7a54a5ee65757af77105","Command":"/bin/run axiom/containers","Created":1485470687,"Ports":[],"Labels":{"com.docker.swarm.node.id":"wa89k0ko2e1qfjc0mana1m5p8","com.docker.swarm.service.id":"v1ikpy2p6mwl9sjcdt4llpwtn","com.docker.swarm.service.name":"shell","com.docker.swarm.task":"","com.docker.swarm.task.id":"o6zf332yjf3idld9y3go7aell","com.docker.swarm.task.name":"shell.1.o6zf332yjf3idld9y3go7aell"},"State":"running","Status":"Up About a minute","HostConfig":{"NetworkMode":"default"},"NetworkSettings":{"Networks":{"axiom":{"IPAMConfig":{"IPv4Address":"10.0.0.5"},"Links":null,"Aliases":null,"NetworkID":"lhx2u89p26rtaqg4aurexi3y3","EndpointID":"f19f1060b945fb1d51556abc82b29927049db1ccdc3805a60c7f9c8dbc1ea72d","Gateway":"","IPAddress":"10.0.0.5","IPPrefixLen":24,"IPv6Gateway":"","GlobalIPv6Address":"","GlobalIPv6PrefixLen":0,"MacAddress":"02:42:0a:00:00:05"}}},"Mounts":[]},{"Id":"70dedbd2cc469721c74bbae3001e4d5923066b3d524b8071fadda69e0e914a32","Names":["/axiom.1.sby4vhnwtn9a8tu16rj7jx86q"],"Image":"ehazlett/axiom@sha256:018046342f7894d1c4eb123740ff0d52861d697550c46ae04978b42f66502052","ImageID":"sha256:eca54ad09c14f50a69ba4ac04577f8f6d0b6eddff73fc3ebc88261ed0b05b913","Command":"/bin/app -D run","Created":1485470598,"Ports":[{"PrivatePort":8080,"Type":"tcp"}],"Labels":{"com.docker.swarm.node.id":"wa89k0ko2e1qfjc0mana1m5p8","com.docker.swarm.service.id":"xobg1a5oi998z8bwmpo9mj33r","com.docker.swarm.service.name":"axiom","com.docker.swarm.task":"","com.docker.swarm.task.id":"sby4vhnwtn9a8tu16rj7jx86q","com.docker.swarm.task.name":"axiom.1.sby4vhnwtn9a8tu16rj7jx86q"},"State":"running","Status":"Up 3 minutes","HostConfig":{"NetworkMode":"default"},"NetworkSettings":{"Networks":{"axiom":{"IPAMConfig":{"IPv4Address":"10.0.0.3"},"Links":null,"Aliases":null,"NetworkID":"lhx2u89p26rtaqg4aurexi3y3","EndpointID":"bfaa83eace7c51f015569eeee75775665cf67bcd7ee62433473c1cff30cfea2d","Gateway":"","IPAddress":"10.0.0.3","IPPrefixLen":24,"IPv6Gateway":"","GlobalIPv6Address":"","GlobalIPv6PrefixLen":0,"MacAddress":"02:42:0a:00:00:03"}}},"Mounts":[{"Type":"bind","Source":"/var/run/docker.sock","Destination":"/var/run/docker.sock","Mode":"","RW":false,"Propagation":""}]}]
```

Note: Axiom currently only provides readonly access.  It is meant as an
information service not a control point.

# Scope
Axiom supports scoping to limit access.  To enable, create the service with
the `--scope=limited` option:

```
$> docker service create \
    --name axiom \
    --network axiom \
    --mount type=bind,src=/var/run/docker.sock,target=/var/run/docker.sock,ro=1 \
    ehazlett/axiom run --scope=limited
```

This will only output information about the container that makes the request.

Currently this is limited to containers and services but will be extended
to limit scope to other resources such as stacks and networks.
