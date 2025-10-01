Package ```customerimporter``` reads from a CSV file and returns a sorted (data
structure of your choice) of email domains along with the number of customers
with e-mail addresses for each domain. 

* Initialise a new git repository and commit your changes.
* Fix/Extend the project so that it runs from the CLI and output the sorted domains to the terminal or to a file. 
* Any errors should be logged and handled.
* Tests should pass.
* Performance matters (this is only ~3k lines, but could be 1m lines or run on a small machine).
* You are free to refactor/create anything the demonstrates how you build code.


# What was done
## Customer importer
* Left error handling as is, because strict validation is better when importing customer data.
* Left email validation as is, since it is good enough for this task and i focused on performance first.
* Moved DomainData from exporter to the entity - decoupled the exporter from the customer importer.
* Did some restructuring: moved email validation and sorting into separate functions.
* Added ReuseRecord for CSV reading - decreased memory usage.
* Changed the sorting algorithm: sort by simple slice domain first, then create a sorted slice with structures - performance optimization, since sorting by domains in the simple slice is faster.

## Exporter
* Created an interface to allow adding more exporters in the future.
* Divided exporter into CSV and terminal implementations.
* The main function now gets an exporter using GetExporter - makes the code cleaner and easier to read.
* Minor change for the pair variable – initialized it before the loop.

## Notes

I do not know if goroutines are expected as part of this task, but there is no sense in using them here, because importer parts such as email validation are very fast. Using goroutines could actually decrease performance due to context switching and locking overhead.
