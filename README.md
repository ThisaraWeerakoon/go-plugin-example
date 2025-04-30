# go-plugin-example
This repo provides a simplified example of how to support plugins in go. Future contributors can can use this as a PoC and continue on this approach for implementing a connectore store for Apache Synpase Go implementation. I wrote a [blog article](https://medium.com/@thisara.weerakoon2001/building-extensible-go-applications-with-plugins-19a4241f3e9a) also on descirbing the concepts behind this simple demo. 

## Component Diagram

This diagram shows the high-level packages/components and their dependencies.

```mermaid
graph TD
    subgraph "Main Application"
        MainApp[cmd/main]
    end

    subgraph "Protocol Definition"
        Protocol[protocol]:::protocolStyle
        Protocol -- defines --> LogPluginInterface(LogPlugin Interface)
    end

    subgraph "Plugins (.so files)"
        SimplePlugin[simple-log-plugin]:::pluginStyle
        FancyPlugin[fancy-log-plugin]:::pluginStyle
    end

    MainApp -- depends on --> Protocol
    MainApp -- dynamically loads --> SimplePlugin
    MainApp -- dynamically loads --> FancyPlugin

    SimplePlugin -- depends on --> Protocol
    FancyPlugin -- depends on --> Protocol

    SimplePlugin -- implements --> LogPluginInterface
    FancyPlugin -- implements --> LogPluginInterface

    classDef protocolStyle fill:#e0f2fe,stroke:#38bdf8,stroke-width:2px;
    classDef pluginStyle fill:#fef3c7,stroke:#f59e0b,stroke-width:2px;
```

## Sequence Diagram

This diagram shows the sequence of interactions when the main application loads and uses a plugin.

```mermaid
sequenceDiagram
    participant Main as MainApp
    participant PluginLib as plugin (Go Lib)
    participant SO as LoadedPluginSO
    participant Symbol as PluginSymbol
    participant LogPlugin as SpecificPlugin:LogPlugin

    Main->>PluginLib: Open(pluginPath)
    activate PluginLib
    PluginLib-->>Main: SO (Plugin Reference)
    deactivate PluginLib

    Main->>SO: Lookup("Plugin")
    activate SO
    SO-->>Main: Symbol (interface{})
    deactivate SO

    Main->>Main: Assert: logPlugin = Symbol.(LogPlugin)

    alt Loading Successful
        Main->>LogPlugin: PrintMessage("Hello Plugin!")
        activate LogPlugin
        LogPlugin-->>Main: (Logging action occurs)
        deactivate LogPlugin
    else Loading Failed (Open, Lookup, or Assert)
        Main->>Main: Handle Error
    end
```
