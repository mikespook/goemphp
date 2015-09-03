#include "zend.h"
void php_set_ini(char *ini);
void php_startup();
char * php_exec_file(char *filename);
char * php_eval_script(char *script);
void php_shutdown(void);
void php_add_var_str(char *varname, char *value);
void php_add_var_double(char *varname, double value);
void php_add_var_long(char *varname, long value);
void php_add_var_bool(char *varname, int value);
zval * php_array_init();
void php_array_add(zval *arr, char *key, char *value);
void php_array_end(zval *arr, char *varname);
void php_unset(char *varname);
