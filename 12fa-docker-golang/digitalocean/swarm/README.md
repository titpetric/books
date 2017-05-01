# Auto-scaling a Docker swarm

The set of scripts here create, grow or shrink a Docker Swarm by provisioning droplet
instances on Digital Ocean. If you haven't signed up yet, you can [use this affiliate
link to give you $10 in credit](https://m.do.co/c/f6a67e30a1f4).

This requires that you have [doctl](https://github.com/digitalocean/doctl) installed.

The scripts are produced for:

1. [12 Factor Applications with Docker and Go](https://leanpub.com/12fa-docker-golang)
2. [Upcoming article on scene-si.org](https://scene-si.org)

## Nodes

There are two types of nodes that exist in a docker swarm.

### Managers

Manager nodes handle the main orchestration part of the swarm cluster. You can
have as many manager nodes, but you need a majority to function correctly. This
means if you have 3 managers, you can remove one without an outage. If you have
five, you can remove 2 without an outage.

To add a manager node, run:

```
./add-manager.sh
```

This will take about 2 minutes until the first swarm is created. After that you
can add additional managers or worker nodes at a higher speed / concurrently.
The join token is being passed to the instance over a cloud-init script.

### Workers

Worker nodes don't get a management interface to the swarm, but just an execution.
This means that you can't perform any swarm actions from the workers, but the
managers schedule containers for execution on them. They work, not manage.

To add a worker node, run:

```
./add-worker.sh
```

You can have as many workers as you like, you can even have none if you don't.
The join token is being passed to the instance over a cloud-init script.

### Removing managers

Removing a manager first demotes the manager to a worker, and then the worker
node is removed from the swarm, after which, the digital ocean droplet is purged.

So:

1. Demote manager to worker,
2. Remove worker node from swarm,
3. Purge droplet

All this can be done by running:

```
./remove-manager.sh
```

For example:

~~~
# ./list-swarm.sh
ID                           HOSTNAME          STATUS  AVAILABILITY  MANAGER STATUS
uoks2o8ce27dl0w1iz9upd3xz    swarm-1493615119  Ready   Active        Leader
x1njngdyj53tbhxjtlk8ie4fh    swarm-1493615341  Ready   Active        Reachable
ybp04oaayxhv82agvfrfrraja *  swarm-1493615112  Ready   Active        Reachable
# ./remove-manager.sh
Leaving swarm: swarm-1493615112
Manager swarm-1493615112 demoted in the swarm.
Purging droplet: swarm-1493615112
swarm-1493615112
# ./list-swarm.sh
ID                           HOSTNAME          STATUS  AVAILABILITY  MANAGER STATUS
uoks2o8ce27dl0w1iz9upd3xz *  swarm-1493615119  Ready   Active        Leader
x1njngdyj53tbhxjtlk8ie4fh    swarm-1493615341  Ready   Active        Reachable
~~~

As the node is removed gracefully, the constraints about fault tolerance do not apply
any more. The swarm size goes down, and with it, availability guarantees and constraints.

Before the manager is removed, it's availability is set to drain mode for a few seconds.
This should be enough for the swarm to start re-scheduling any containers to other nodes.

### Removing workers

Removing a worker is just as simple as removing a manager:

~~~
./remove-worker.sh
~~~

You can remove any number of workers without causing an outage, as long as your
managers have enough capacity to start all the containers.

### Destroying everything

Just run `./destroy.sh` and watch the world burn.

### Other

There are a few utility scripts to facilitate the functionality presented above:

* `./list-managers.sh` - lists running manager nodes,
* `./list-workers.sh` - list running worker nodes,
* `./list-swarm.sh` - prints output of `docker node ls`, showing nodes in swarm,
* `./list.sh` - show all digital ocean droplets running,
* `./ssh-key.sh` - provide ssh key to digital ocean instance for logging in,
* `./ssh.sh` - run a command on all manager nodes

For example, if you want to run `uname` on all manager nodes, you can do:

~~~
# ./ssh uname -r
> swarm-1493616101
4.4.0-75-generic
> swarm-1493616219
4.4.0-75-generic
> swarm-1493616226
4.4.0-75-generic
~~~

### Testing the swarm

After spinning up the managers, you can create a service on them:

~~~
# ./ssh-one.sh docker service create --replicas 10 --name sonyflake titpetric/sonyflake
> swarm-1493616101
ld66ax987n5nyypm1itegy6io
~~~

~~~
# ./ssh-one.sh docker service ps sonyflake --format '{{.Node}}' \| sort \| uniq -c
> swarm-1493616101
      4 swarm-1493616101
      3 swarm-1493616219
      3 swarm-1493616226
~~~

And removing a manager:

~~~
# ./remove-manager.sh
Leaving swarm: swarm-1493616101
swarm-1493616101
Manager swarm-1493616101 demoted in the swarm.
Purging droplet: swarm-1493616101
swarm-1493616101
~~~

~~~
# ./ssh-one.sh docker service ps sonyflake --format '{{.Node}}' -f 'desired-state=running' \| sort \| uniq -c
> swarm-1493616219
      5 swarm-1493616219
      5 swarm-1493616226
~~~

Here we see that the containers have re-scheduled on the available managers without issue. As long
as there is one manager left, we can remove managers without failure of the swarm.

~~~
# ./remove-manager.sh
Leaving swarm: swarm-1493616219
swarm-1493616219
Manager swarm-1493616219 demoted in the swarm.
Purging droplet: swarm-1493616219
swarm-1493616219
~~~

~~~
# ./ssh-one.sh docker service ps sonyflake --format '{{.Node}}' -f 'desired-state=running' \| sort \| uniq -c
> swarm-1493616226
     10 swarm-1493616226
# ./list-swarm.sh
ID                           HOSTNAME          STATUS  AVAILABILITY  MANAGER STATUS
naoyvnd0nu1vxegi0mplxy1sf *  swarm-1493616226  Ready   Active        Leader
~~~

### Conclusion

With the set of these scripts it's possible to provide swarm elasticity. If your load is elastic,
you can use the provided scripts to add and remove worker and manager nodes as needed to faster
process your workloads. Depending on any monitoring rules you set up, you can create your own system
to grow and scale your docker swarm based on a number of inputs that you monitor. This may be anything
from CPU usage, to the number of items in your worker queue (or other application logic).
