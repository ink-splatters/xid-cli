# xid-cli

Generation of globally unique, monotonous identifiers.

CLI for [xid](https://github.com/rs/xid)

## Usage

Just type

```sh
xid
```

for generation of new xid

Info about existing xid:

```sh
xid -info <xid>
```

xids are time dependent, among other things, and can be created from unix timestamp as well:

```sh
xid -timestamp <unix timestamp>
```
