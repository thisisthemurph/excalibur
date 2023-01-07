using Microsoft.AspNetCore.Http;

namespace Excalibur.Application.Services;
public interface IFileUploadService
{
    Task StoreFile(IFormFile formFile, string newFileName);
}