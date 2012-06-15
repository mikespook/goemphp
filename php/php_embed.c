#include "php_embed.h"

/* {{{ goemphp php_startup(ini)
*/
void php_startup(char *ini) {
    if (php_embed_module.php_ini_path_override) {
        free(php_embed_module.php_ini_path_override);
    }
    php_embed_module.php_ini_path_override = strdup(ini);
    php_embed_init(0, NULL PTSRMLS_CC);
}
/* }}} */

/* {{{ goemphp php_exec_file(filename)
*/
char * php_exec_file(char *filename) {
    char *result = NULL;
    zend_first_try {
        zend_file_handle file_handle;
        file_handle.type = ZEND_HANDLE_FILENAME;
        file_handle.filename = filename;
        file_handle.free_filename = 0;
        file_handle.opened_path = NULL;
        if (php_execute_script( &file_handle TSRMLS_CC ) == FAILURE) {
            if (PG(last_error_message)) {
                result = strdup(PG(last_error_message));
                free(PG(last_error_message));
                PG(last_error_message) = NULL;
            }
        }
    } zend_end_try();
    return result;
}
/* }}} */

char * php_eval_script(char *script) {
    char *result = NULL;
    zend_first_try {
        if ( zend_eval_string(script, NULL, "GoEmPHP" TSRMLS_CC) == FAILURE ) {
            if (PG(last_error_message)) {
                result = strdup(PG(last_error_message));
                free(PG(last_error_message));
                PG(last_error_message) = NULL;
            }
        }
    } zend_end_try();
    return result;
}

void php_shutdown(void) {
    php_embed_shutdown(TSRMLS_CC);
}

int php_info(void) {
    if (php_request_startup(TSRMLS_C)==FAILURE) {
        return FAILURE;
    }
    php_print_info(0xFFFFFFFF TSRMLS_CC);
    php_end_ob_buffers(1 TSRMLS_CC);
    return SUCCESS;
}
