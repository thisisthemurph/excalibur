using Excalibur.Api.Exceptions;

namespace Excalibur.Api.Services;

public class FileUploadService
{
    private readonly ILogger<FileUploadService> _logger;
    private readonly IWebHostEnvironment _hostEnvironment;

    public FileUploadService(
        IWebHostEnvironment hostEnvironment, 
        ILogger<FileUploadService> logger)
    {
        _logger = logger ?? throw new ArgumentNullException(nameof(logger));
        _hostEnvironment = hostEnvironment ?? throw new ArgumentNullException(nameof(hostEnvironment));
    }

    public async Task StoreFile(IFormFile formFile, string newFileName)
    {
        string uploadsPath = Path.Combine(_hostEnvironment.WebRootPath, "uploads");

        // Validate

        if (formFile is null || formFile.Length == 0)
        {
            _logger.LogError("Uploaded file is null or blank {FileName}", newFileName);
            throw new EmptyFileException("Expected file content, blank or null reference file found.");
        }

        // Ensure the storage location

        string filePath = Path.Combine(uploadsPath, newFileName);
        var directory = Path.GetDirectoryName(filePath);
        if (directory is null) 
        {
            _logger.LogError("Directory not discernible from file path: {FilePath}.", filePath);
            throw new DirectoryNotFoundException("Could not deternmine the directory storage location.");    
        }

        Directory.CreateDirectory(directory);

        // Upload / copy file

        using var stream = new FileStream(filePath, FileMode.Create);
        await formFile.CopyToAsync(stream);
    }
}
