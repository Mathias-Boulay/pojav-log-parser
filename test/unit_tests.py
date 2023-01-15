import unittest
import main


class TestParserMethods(unittest.TestCase):
    content_log = ''

    def setUp(self) -> None:
        self.content_log = open('../logs/latestlog_0.txt', 'rt').read()




if __name__ == '__main__':
    unittest.main()
