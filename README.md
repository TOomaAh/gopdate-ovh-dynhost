# gopdate-ovh-dynhost


gopdate is a script to update OVH dynhost with a readable and easy to use configuration file.

You can update the configuration file even if the script is running and the new configuration will be taken over at the next iteration.

## Usage

###Configuration example

If no configuration file is found a new one will be created

configuration:
```
{
    "seconds":300,
    "config":[
        {
            "domain": "",
            "username": "",
            "pasword": ""
        }
    ]
}

```
