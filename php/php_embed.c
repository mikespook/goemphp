// Copyright 2011 Xing Xing <mikespook@gmail.com> All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

#include "sapi/embed/php_embed.h"
#include "zend.h"
#include <unistd.h>

/* This function doesn't return if it uses E_ERROR */
zend_class_entry *default_exception_ce;
char * exception_error(zval *exception, int severity TSRMLS_DC) {
    char * result;
    zend_class_entry *ce_exception = Z_OBJCE_P(exception);
    if (instanceof_function(ce_exception, default_exception_ce TSRMLS_CC)) {
        zval *str, *file, *line;

        EG(exception) = NULL;
        zend_call_method(&exception, ce_exception, NULL, "__tostring", 10, &str, 0, NULL, NULL TSRMLS_CC);
        if (!EG(exception)) {
            if (Z_TYPE_P(str) != IS_STRING) {
                result = "Exception::__toString() must return a string";
            } else {
                zend_update_property_string(default_exception_ce, exception, "string", sizeof("string")-1, EG(exception) ? ce_exception->name : Z_STRVAL_P(str) TSRMLS_CC);
            }
        }
        zval_ptr_dtor(&str);

        if (EG(exception)) {
            /* do the best we can to inform about the inner exception */
            if (instanceof_function(ce_exception, default_exception_ce TSRMLS_CC)) {
                file = zend_read_property(default_exception_ce, EG(exception), "file", sizeof("file")-1, 1 TSRMLS_CC);
                line = zend_read_property(default_exception_ce, EG(exception), "line", sizeof("line")-1, 1 TSRMLS_CC);
            } else {
                file = NULL;
                line = NULL;
            }
            result = "Uncaught in exception handling during call to __tostring()";
        }

        str = zend_read_property(default_exception_ce, exception, "string", sizeof("string")-1, 1 TSRMLS_CC);
        file = zend_read_property(default_exception_ce, exception, "file", sizeof("file")-1, 1 TSRMLS_CC);
        line = zend_read_property(default_exception_ce, exception, "line", sizeof("line")-1, 1 TSRMLS_CC);

        result = "Uncaught %s\n  thrown";
    } else {
        size_t n = snprintf(NULL, 0, "Uncaught exception '%s'", ce_exception->name);
        result = malloc(n + 1);
        snprintf(result, n + 1, "Uncaught exception '%s'", ce_exception->name);
    }
    return result;
}
/* }}} */


/* {{{ goemphp php_set_ini(ini)
*/
void php_set_ini(char *ini) {
    if (php_embed_module.php_ini_path_override) {
        free(php_embed_module.php_ini_path_override);
    }
    php_embed_module.php_ini_path_override = strdup(ini);
}
/* }}} */

/* {{{ goemphp php_startup(ini)
*/
void php_startup() {
    php_embed_init(0, NULL PTSRMLS_CC);
}
/* }}} */

/* {{{ goemphp php_eval_script(script) 
*/
char * php_eval_script(char *script) {
    char *result = NULL;
    zend_first_try {
        zend_eval_string(script, NULL, "GoEmPHP" TSRMLS_CC);
        // executed failure, get error message
        if (PG(last_error_message)) {
            result = strdup(PG(last_error_message));
            free(PG(last_error_message));
            PG(last_error_message) = NULL;
        }

        if (PG(last_error_message)) {
            result = strdup(PG(last_error_message));
            free(PG(last_error_message));
            PG(last_error_message) = NULL;
        }

        if (!result && EG(exception)) {
            result = exception_error(EG(exception), E_ERROR TSRMLS_CC);
            EG(exception) = NULL;
        }
        // trigger_error & throw exception need to handle
        // if (error) {
        //
        // } else if (exception) {
        //
        // }
    } zend_end_try();
    return result;
}
/* }}} */

/* {{{ goemphp php_exec_file(filename)
*/
char * php_exec_file(char *filename) {
    char *result = NULL;
    zend_first_try {
        zend_file_handle file_handle = {0};
        file_handle.type = ZEND_HANDLE_FILENAME;
        file_handle.filename = filename;
        file_handle.free_filename = 0;
        file_handle.opened_path = NULL;

        php_execute_script(&file_handle TSRMLS_CC ); 
        // executed failure, get error message
        if (PG(last_error_message)) {
            result = strdup(PG(last_error_message));
            free(PG(last_error_message));
            PG(last_error_message) = NULL;
        }
        if (!result && EG(exception)) {
            result = exception_error(EG(exception), E_ERROR TSRMLS_CC);
            EG(exception) = NULL;
        }
        // trigger_error & throw exception need to handle
        // if (error) {
        //
        // } else if (exception) {
        //
        // } 
    } zend_end_try();
    return result;
}
/* }}} */

/* {{{ goemphp php_shutdown()
*/
void php_shutdown(void) {
    php_embed_shutdown(TSRMLS_CC);
}
/* }}} */

/* {{{ goemphp php_add_var_bool(char *varname, int value)
 */
void php_add_var_bool(char *varname, int value) {
    zval *newvar;
    MAKE_STD_ZVAL(newvar);
    ZVAL_BOOL(newvar, value);
    zend_hash_update(&EG(symbol_table), varname, strlen(varname) + 1, &newvar, sizeof(zval *), NULL);
}
/* }}} */

/* {{{ goemphp php_add_var_long(char *varname, long value)
 */
void php_add_var_long(char *varname, int value) {
    zval *newvar;
    MAKE_STD_ZVAL(newvar);
    ZVAL_LONG(newvar, value);
    zend_hash_update(&EG(symbol_table), varname, strlen(varname) + 1, &newvar, sizeof(zval *), NULL);
}
/* }}} */

/* {{{ goemphp php_add_var_double(char *varname, double value)
 */
void php_add_var_double(char *varname, double value) {
    zval *newvar;
    MAKE_STD_ZVAL(newvar);
    ZVAL_DOUBLE(newvar, value);
    zend_hash_update(&EG(symbol_table), varname, strlen(varname) + 1, &newvar, sizeof(zval *), NULL);
}
/* }}} */

/* {{{ goemphp php_add_var_str(char *varname, char *value)
 */
void php_add_var_str(char *varname, char *value) {
    zval *newvar;
    MAKE_STD_ZVAL(newvar);
    ZVAL_STRING(newvar, value, 1);
    zend_hash_update(&EG(symbol_table), varname, strlen(varname) + 1, &newvar, sizeof(zval *), NULL);
}
/* }}} */


