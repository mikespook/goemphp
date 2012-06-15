#include "php_embed.h"
#ifdef ZTS
    void ***tsrm_ls;
#endif
/* Extension bits */
zend_module_entry goemphp_module_entry = {
    STANDARD_MODULE_HEADER,
    "goemphp", /* extension name */
    NULL, /* function entries */
    NULL, /* MINIT */
    NULL, /* MSHUTDOWN */
    NULL, /* RINIT */
    NULL, /* RSHUTDOWN */
    NULL, /* MINFO */
    "1.0", /* version */
    STANDARD_MODULE_PROPERTIES
};
void php_startup(void) {
    int argc = 1;
    char *argv[2] = { "goemphp", NULL };
    php_embed_init(argc, argv PTSRMLS_CC);
    zend_startup_module(&goemphp_module_entry);
}

void php_exec_file(char *filename) {
    zend_first_try {
        char *include_script;
        spprintf(&include_script, 0, "include '%s'", filename);
        zend_eval_string(include_script, NULL, filename TSRMLS_CC);
        efree(include_script);
    } zend_end_try();
}

void php_eval_script(char *script) {
    zend_first_try {
        zend_eval_string(script, NULL, "GoEmPHP" TSRMLS_CC);
    } zend_end_try();
}

void php_shutdown(void) {
    php_embed_shutdown(TSRMLS_CC);
}
