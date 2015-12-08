# GoVMAX

## Overview
```GoVMAX``` represents API bindings for Go that allow you to manage VMAX
block platforms.  In the true nature of API bindings, it is intended that the
functions available are basically a direct implementation of what is available
through the API.  

The bindings also include VMware vSphere functionality to enable use cases
when mapping VMAX storage to RDMs on VMs.

## API Compatibility
Currently only tested with VMAX3.

## Examples
The package was written using test files, so these can be looked at for a more
comprehensive view of how to implement the different functions.

Intialize a new client

    smis, err = New(host, port, insecure, username, password)


### Some Volume Examples

Get Storage Pools

    pools, err := smis.GetStoragePools(testingSID)


Get Volumes

    vols, err := smis.GetVolumes(testingSID)



For example usage you can see the [REX-Ray](https://github.com/emccode/rexray)
repo.  There, the ```govmax``` package is used to implement a
```Volume Manager``` across multiple storage platforms. This includes managing
multipathing, mounts, and filesystems.

## Environment Variables
Name | Description
---- | -----------
`GOVMAX_SMISHOST` | the API host
`GOVMAX_SMISPORT` | the API host port
`GOVMAX_USERNAME` | the username
`GOVMAX_PASSWORD` | the password
`GOVMAX_INSECURE` | whether to skip SSL validation

## Contributions
Please contribute!

Licensing
---------
Licensed under the Apache License, Version 2.0 (the “License”); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.

Support
-------
If you have questions relating to the project, please either post
[Github Issues](https://github.com/emccode/govmax/issues), join our
Slack channel available by signup through
[community.emc.com](https://community.emccode.com) and post questions into
`#projects`, or reach out to the maintainers directly.  The code and
documentation are released with no warranties or SLAs and are intended to be
supported through a community driven process.
