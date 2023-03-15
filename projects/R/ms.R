install.packages("tidyverse")
library(tidyverse)

glimpse(msleep)

missing <- !complete.cases(msleep)

msleep[missing, ]
