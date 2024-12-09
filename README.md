# Unifi Toolkits

**English** | [中文](./README_zh.md)

Some functionalities are not available on the web console,
neither unifi os nor self-hosted controller.
It causes the inconvenience to use.

For improving the experience to use the web console,
this project provides some tools as follows:

### 1. Prune offline clients
There are lots of guest clients remaining in the client list, which can only be removed one by one.

This tool can do bulk deletion. Only the offline clients without a custom name will be pruned.
It is exactly the same action that you click "Remove" button on the web console.

### 2. Print mac address filter
The clients' name are not displayed on the mac address filter list, making them difficult to be recognized.

This tools will print the list under the corresponding WiFi with mac address and client name.

## Arguments
- -m: Model, required<br>
      Only **Console** or **Controller** can be applied<br>
      Console relates to Unifi OS like UDMPro, UDR<br>
      Controller relates to self-hosted controller
- -g: Console/Controller IP Address, **required**
- -p: Console/Controller Port, **optional**. **Default is 443**
- -u: Console/Controller User Name, **required**
- -d: Dry Run, **optional**. Default is **false**

## Examples
```shell
./print-mac-filter -m Console -g 192.168.10.1 -u xxx

./print-mac-filter -m Console -g 192.168.10.1 -u xxx -d true # run before pruning
./print-mac-filter -m Console -g 192.168.10.1 -u xxx
```