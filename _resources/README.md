# Resources

This folder contains several [resources](https://learn.microsoft.com/en-us/windows/win32/menurc/about-resource-files) that can be used to build your native Win32 application:

| Resource | Description |
| - | - |
| `gopher.ico` | An icon that can be used as default. |
| `win10.exe.manifest` | A basic manifest file that enables your application to be recognized as a Windows 10 one. |
| `minimal.res` | A compiled Win32 resource script, which contains the icon and the manifest. |
| `minimal.syso` | A syso file, ready to use, which contains the icon and the manifest, built from `minimal.syso`. Just place it at the root folder of your project. You can load the icon using the resource ID 101. |

The `.res` file can be created/edited with [Visual Studio](https://visualstudio.microsoft.com/) or [Resoure Hacker](https://www.angusj.com/resourcehacker/).

The `.res` file can be converted into `.syso` with [windres](https://winlibs.com/#download-release):

```
windres.exe -i minimal.res -o minimal.syso
```
