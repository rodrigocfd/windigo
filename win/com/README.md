# Windigo / com

This package contains packages with bindings for specific COM libraries.

Each COM package will always have two subpackages, with the following suffixes:

* `co` for the constants of that COM package;
* `vt` for the virtual tables.

Specifically, the `com/com` package contains the COM foundations like `IUnknown` and `CoCreateInstance`, which are used by all other COM packages.
