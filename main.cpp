#include <iostream>
#include <string>
#include <vector>
#include <boost/filesystem.hpp>
#include <boost/program_options.hpp>
#include <boost/interprocess/file_mapping.hpp>
#include <boost/interprocess/mapped_region.hpp>
using namespace boost::program_options;
using namespace boost::interprocess;

const int N = 128;

// read from input file and split to 128 smaller files, hash(x) % 128
// count in each smaller file
// pick up the Top N items from each smaller file
// pick up the Top N items from 128 * Top N items

bool countItemsInFile(const std::string& filePath)
{
    return true;
}

bool countItemsInAll(const std::vector<std::string>& fileNames)
{
    return true;
}

int hash(const std::string& str)
{
    // inspired by hash algorithme from Java 8 HashMap
    int res = 0;
    for (size_t i = 0; i < str.size(); i++)
    {
        res += 31 * res + str[i];
    }
    res ^= (res >> 16);
    return res & (N-1); // means (res % 128)
}

bool readSplitInput(const std::string& inputFile, const std::vector<std::string>& fileNames)
{
    // use file mapping to accelerate large file reading
    try {
        //Create a file mapping
        file_mapping m_file(inputFile.c_str(), read_only);

        offset_t offset = 0;
        std::size_t page_size = mapped_region::get_page_size();
        //Map the file with read-only permissions in this process
        mapped_region region( m_file                         //Map shared memory
                            , read_only                     //Map it as read-only
                            , offset                        //Map from offset 0
                            , page_size                     //Map until the end
                            );

        //Get the address of the mapped region
        unsigned char * addr = static_cast<unsigned char *>(region.get_address());
        std::size_t size  = region.get_size();


    } catch (std::exception& e) {
        std::cerr << e.what() << std::endl;
        return false;
    }
    return true;
}

int main(int argc, char* argv[])
{
    std::cout << "Pick up Top N items from file.\n";

    std::string inputFile = "output.txt";
    std::string tempDir = "./";
    options_description desc("Allowed options");
    desc.add_options()
    ("help,h", "print usage message")
    ("input,i", value(&inputFile), "path of input data file")
    ("temp,t", value(&tempDir), "temporary directory path")
    ;

    variables_map vm;
    store(parse_command_line(argc, argv, desc), vm);

    if (vm.count("help")) 
    {  
        std::cout << desc << std::endl;
        return 0;
    }

    std::vector<std::string> fileNames;
    fileNames.reserve(N);
    for (int i = 0; i < N; i ++)
        fileNames.push_back(tempDir + "/" + std::to_string(i));

    std::cout << "read from file " << inputFile << ", use " << tempDir << " as temporary directory.\n";
    if (!readSplitInput(inputFile, fileNames))
    {
        return 1;
    }

    for(const auto& filePath : fileNames)
    {
        if (!countItemsInFile(filePath + ".txt"))
            return 2;
    }    

    if (!countItemsInAll(fileNames))
    {
        return 3;
    }

    for(const auto& filePath : fileNames)
    {
        boost::filesystem::remove(boost::filesystem::path(filePath+".txt"));
        boost::filesystem::remove(boost::filesystem::path(filePath+".processed.txt"));
        boost::filesystem::remove(boost::filesystem::path(filePath+".result.txt"));
    }

    return 0;
}