package main

const containerJson string = `
{
    "capabilities": [
        "CHOWN",
        "DAC_OVERRIDE",
        "FOWNER",
        "MKNOD",
        "NET_RAW",
        "SETGID",
        "SETUID",
        "SETFCAP",
        "SETPCAP",
        "NET_BIND_SERVICE",
        "SYS_CHROOT",
        "KILL"
    ],
    "cgroups": {
        "allowed_devices": [
            {
                "cgroup_permissions": "m",
                "major_number": -1,
                "minor_number": -1,
                "type": 99
            },
            {
                "cgroup_permissions": "m",
                "major_number": -1,
                "minor_number": -1,
                "type": 98
            },
            {
                "cgroup_permissions": "rwm",
                "major_number": 5,
                "minor_number": 1,
                "path": "/dev/console",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "major_number": 4,
                "path": "/dev/tty0",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "major_number": 4,
                "minor_number": 1,
                "path": "/dev/tty1",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "major_number": 136,
                "minor_number": -1,
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "major_number": 5,
                "minor_number": 2,
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "major_number": 10,
                "minor_number": 200,
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 3,
                "path": "/dev/null",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 5,
                "path": "/dev/zero",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 7,
                "path": "/dev/full",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 5,
                "path": "/dev/tty",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 9,
                "path": "/dev/urandom",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 8,
                "path": "/dev/random",
                "type": 99
            }
        ],
        "name": "docker-koye",
        "parent": "docker"
    },
    "restrict_sys": true,
    "mount_config": {
        "device_nodes": [
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 3,
                "path": "/dev/null",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 5,
                "path": "/dev/zero",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 7,
                "path": "/dev/full",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 5,
                "path": "/dev/tty",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 9,
                "path": "/dev/urandom",
                "type": 99
            },
            {
                "cgroup_permissions": "rwm",
                "file_mode": 438,
                "major_number": 1,
                "minor_number": 8,
                "path": "/dev/random",
                "type": 99
            }
        ],
        "mounts": [
            {
                "type": "tmpfs",
                "destination": "/tmp"
            }
        ]
    },
    "environment": [
        "HOME=/",
        "PATH=/usr/local/bin",
        "HOSTNAME=redis",
        "TERM=xterm"
    ],
    "hostname": "redis",
    "namespaces": {
        "NEWIPC": true,
        "NEWNET": false,
        "NEWNS": true,
        "NEWPID": true,
        "NEWUTS": true
    },
    "tty": false,
    "user": "root"
}
`
