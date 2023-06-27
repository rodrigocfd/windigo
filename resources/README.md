# Resources

This folder contains several [resources](https://learn.microsoft.com/en-us/windows/win32/menurc/about-resource-files) that can be used to build your native Win32 application:

| Resource | Description |
| - | - |
| `gopher.ico` | An icon that can be used as default. |
| `win10.exe.manifest` | A basic manifest file that enables your application to be recognized as a Windows 10 one. |
| `minimal.syso` | A syso file, ready to use, that contains the icon and the manifest. Just place it at the root folder of your project. You can load the icon using the resource ID 101. |

If you wish, you can build your own syso:

* with the [rsrc](https://github.com/akavel/rsrc) tool;
* creating a `.rc` file from scratch and using a resource compiler, like [MSVC/RC](https://learn.microsoft.com/en-us/windows/win32/menurc/resource-compiler).
