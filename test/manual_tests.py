import main as parser


file = open('../logs/latestlog_0.txt', 'rt')
log_content = file.read()

print(log_content)

test = parser.parse_version(log_content)
print(test)

test2 = parser.parse_java_arguments(log_content)
print(test2)

test3 = parser.parse_env_variables(log_content)
print(test3)

test4 = parser.parse_java_runtime(test3['JAVA_HOME'])
print(test4)

test5 = parser.parse_renderer(log_content)
print(test5)

