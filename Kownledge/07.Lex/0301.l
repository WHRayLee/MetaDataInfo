%{
/*
* The word counter example for multiple files
*/
unsigned long charCount = 0, wordCount = 0, lineCount = 0;
#undef yywrap
%}

word [^ \t\n]+
eol \n
%%
{word} {wordCount++;charCount+=yyleng;}
{eol}  {charCount++;charCount++;}
.      {charCount++;}
%%
char **fileList;
unsigned currentFile = 0;
unsigned nFiles = 0;
unsigned long totalCC = 0;
unsigned long totalWC = 0;
unsigned long totalLC = 0;

main(argc, argv)
int argc;
char **argv;
{
    file *file;
    fileList = argv+1;
    nFiles = argc-1;
    if (argc == 2) {
    /*
     * we handle the single file case differently from
     * the multiple file case since we don't need to
     * print a summary line
     */
     currentFile = 1;
     file = fopen(argv[1], "r");
     if (!file) {
     fprintf(stderr,"could not open %s\n",argv[1])
     exit(1);
     }
     yyin = file;
     }
     if (argc > 2)
     yywrap(); /* open first file */
     yylex();
     /*
     * once again, we handle zero or one file
     * differently from multiple files.
     */
     if (argc > 2) {
     printf("%8lu %8lu %8lu %s\n", lineCount, wordCount,
     charCount, fileList[currentFile−1]);
     totalCC += charCount;
     totalWC += wordCount;
      totalLC += lineCount;
      printf("%8lu %8lu %8lu total\n",totalLC, totalWC, totalCC);
    }
}

















