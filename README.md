# SoftwareThatMatters
This document contains the code that was used during the [Research Project](https://github.com/TU-Delft-CSE/Research-Project) (2022 edition) of [TU Delft](https://github.com/TU-Delft-CSE).
## Instructions
This document will help one reproduce the results mentioned in Analyzing the Criticality of Debian Application Through a Time-Dependent Dependency Graph

To set up go dependencies, run the following in the root directory (optional, since go should download deps automatically) :
```
go mod download
```

To run the graph algorithm, simply write in the terminal:
```
go run main.go start
```
This will open up a cli where various commands can be used.

### Note:
To first gather the required data, you need to run the shell script 
``getDebianApss.sh`` from ``data/scripts``. In order to run the script
a Debian-based Linux operating system is required, as the script 
uses ``apt`` commands. The script will create a list with the name
of Debian packages which contain applications. In case you want to
use older package meta-data, you will need to update the ``sources.list``
file in ``/etc/apt`` to also include the ``deb http://deb.debian.org/debian/ oldstable main`` 
repository.

Besides, a JSON file with Debian meta-data is needed in order to build the graph.
The JSON file needs to be in the `/data/input` directory. 
The file can be downloaded from 
### License
The code's main license can be found in LICENSE.
It also re-uses some modified gonum code, for which the license can be found in GONUM_LICENSE