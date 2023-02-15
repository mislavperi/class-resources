###### Data problems
- Rarely perfect data
- Missing data, data errors, coding inconsistencies, missing or bad metadata
- Various reasons for missing data, requires manual labor to cleanup
- Types of missing values vary in existence
	- Null or system missing values
	- Empty strings and white spaces
	- Blank or user defined missing values
###### Handling missing vs noisy data
- Handling noisy
	- Various filters
	- Scatter plots can be used to identify outliers
- Handling missing
	- Missing values can be ignored or filled manually
	- Considered to be bad
	- No best solution, case based
	- Consider
		- Size of data set
		- Number of fields containing blanks
		- Amount of missing information
	- Two approaches generally
		1. Exclude fields with missing values
		2. Impute, replace or coerce missing values
###### Data profiling
- Statistical analysis of values
- Implementation is made shorter with good profiling
- Gives complete and accurate picture of data
- Dataset size, distribution, quality, distinct values, median, mean, quartiles, average, standard deviation
- Percentage of valid, mismatched and empty values
- Number of columns and rows
### Imputation methods
- Form of data cleaning, purpose to fill variables that contain missing values with some intuitive data
- Better to have a reasonable estimate better than leaving blank
- Missing values can be determined through association with rest of data
- Data is not isolated
- Simple approach
	- Not taking account of missingness mechanism
	- Missing at random
		- Data missing at random
	- Missing completely at random
		- Missing value on attribute does not depend on observed or unobserved data
	- Simple methods
		- Identify if column values are required
		- Insert a constant value
		- Use a function
		- Copy values from another column
		- Delete rows
		- Drop the column
	- Inserting constants for missing values
		- Not recommended
		- Shouldn't replace missing numeric values
		- Use only if entire data chain is aware of constants
	- SQL NULL functions
		- `IFNULL(), ISNULL(), COALESCE(), NVL()` functions
		- `IFNULL()` returns alternative value if expression is null
		- `COALESCE()` drop in replacement for `IFNULL()`
	- Machine Learning based Imputation methods
		- Modeling can predict continuous or nominal value from previous learning processes