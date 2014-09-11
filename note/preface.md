1. Go语言支持类，类成员方法，类的组合，但反对继承，反对虚函数和虚函数重载。确切地说，Go也提供了继承，只不过是采用组合的文法来提供。
2. Go语言放弃了构造函数和析构函数。由于Go语言中没有虚函数，也就没有vptr，支持构造函数和析构函数就没有太大的价值。
3. 每个Go源代码文件的开头都是一个package声明，表示该Go代码所属的包。包是Go语言里最基本的分发单位，也是工程管理中依赖关系的体现。要生成Go可执行程序，必须建立一个名字为main的包，并且在该包中包含一个叫main()的函数.