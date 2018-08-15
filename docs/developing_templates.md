# Developing Templates

Ironman provides a way to define templates using a metadata file named ***.ironman.yaml***. A template includes one or more generators. Each generator can include one or more template files which will be processed using parameters passed by the user when the ```ironman generate``` command is called. Files are processed using the [Go templating language](https://golang.org/pkg/text/template/). 

With template generators you can:

***Create a single file***

```
ironman generate my-template:file-generator-name filename --set par1=val1,par2=val2
```

***Create a directory with multiple files inside***

```
ironman generate my-template-file:directory-generator-name directory-name --set par1=val1,par2=val2
```

## Creating a base template

Create a base template using the ```ironman create /path/to/new/template-name```.

That will create the following file structure:  

    
    ├── .ironman.yaml
    ├── README.md
    └── generators
        ├── app
        │   ├── .ironman.yaml
        │   └── README.md
        └── single
            ├── .ironman.yaml
            └── file.txt

The generated base template contains the following.

* Root ***.ironman.yaml***
* Generators directory. Each generator directory includes an ***.ironman.yaml*** file.

## Template

The template is a directory and it must include a ***.ironman.yaml*** file and a ***generators directory***, which will include a directory per generator. Each generator directory also includes an ***.ironman.yaml*** a file at its root. The ***ironman.yaml*** file contains the metadata of templates and generators. 

### Template .ironman.yaml

The .ironman.yaml files contain the information that defines the template.

#### Ironman.yaml properties

* id(mandatory): A unique identifier for the template.
* version(mandatory): A version for the template.
* name(mandatory): A human-readable name for the template.
* description(mandatory): A description for the template.
* home: A URL with a page containing additional information about the template or organization.
* sources: a list of sources for the template.
* maintainers: a list of maintainers for the template.
* deprecated: whether this template should be deprecated.

## Generator

A template contains one or more generators directories. The ***app*** generator is the default generator and will be used when the generate command is called without a generator selector. 

For example, if you want to  to select an specific generator you would run: 

```
ironman generate <template_name>[:generator_name] [path]
```

Or with an explicit example:

```
ironman generate my-template:my-generator /path/to/generated/files
```

With an existing default ***app*** you don't need to select a generator, app generator will be called:

```
ironman generate my-template /path/to/generated/files 
```

### Generator .ironman.yaml

The .ironman.yaml files contain the information that defines the generator.

#### .ironman.yaml properties

* id: Unique identifier for the generator. If not set, it will be infered from the directory name.
* name(mandatory): A human readable name for the generator
* description(mandatory):A description for the generator
* type: The generator type (file | directory)
* fileOptions: Options for the ***file type*** generator.


### Generator types

There are two types of generators.

* ***directory***: The processed output will be all the files in the generator directory.

* ***file***: The processed output will be a single file inside the processed output.

### Generator template file

A generator template file is a file which is going to be rendered based on the user passed parameters when running the generate command. 

A template file should use  [Golang's template language](https://golang.org/pkg/text/template/). The user passed parameters are available in the global "Values" template property.

This is an example on how to use a parameter named "title":

When a generate command is run using a file generator:

    ironman generate mytemplate:mygenerator README.md --set "title=mytitle,subtitle=mysubtitle"

The template file inside de file generator would look something like this:

***README.md file template example***
```md
# {{.Values.title | default "Default title"}}
## {{.Values.subtitle | default "Default sub title" }}
```