using Excalibur.Application.DTOs.Requests;
using Excalibur.Application.DTOs.Responses;
using Excalibur.Application.Repositories;
using Excalibur.Application.Services;
using Excalibur.Domain.Entities;
using Excalibur.Domain.Exceptions;
using Excalibur.Domain.ExtensionMethods;
using Excalibur.Domain.Models;
using Microsoft.AspNetCore.Mvc;

namespace Excalibur.Api.Controllers;

[ApiController]
[Route("[controller]")]
public class FileController : ControllerBase
{
    private readonly ILogger<DataTemplateController> _logger;
    private readonly IDataTemplateRepo _dataTemplateRepo;
    private readonly IFileUploadService _fileUploadService;

    public FileController(
        ILogger<DataTemplateController> logger,
        IDataTemplateRepo dataTemplateRepo,
        IFileUploadService fileUploadService)
    {
        _logger = logger;
        _dataTemplateRepo = dataTemplateRepo;
        _fileUploadService = fileUploadService;
    }

    [HttpPost("upload/{dataTemplateId}")]
    public async Task<IActionResult> UploadSingleFile(string dataTemplateId, IFormFile formFile)
    {
        var ext = Path.GetExtension(formFile.FileName);
        var uniqueDateTime = DateTime.UtcNow.ToIsoFormatString().Replace(":", "-");
        var newFileName = $"UploadedFile_{uniqueDateTime}{ext}";

        // Validate the file name

        if (ext is null)
        {
            return BadRequest("The uploaded file must have a file extension.");
        }

        if (ext.ToUpper() != ".CSV")
        {
            return BadRequest("Only files with a .csv extension are accepted.");
        }

        // Insert the file metadata inton the database

        var updatedDataTemplate = await _dataTemplateRepo.AddFileMetadata(
            dataTemplateId,
            new DataTemplateAddFileMetadataRequest
            {
                Name = formFile.FileName,
                StoredName = newFileName,
                Status = FileUploadStatus.Uploading,
            });

        var newFile = updatedDataTemplate.Files.LastOrDefault();

        if (newFile is null || newFile.Id is null)
        {
            return BadRequest("There has been an issue inserting the file metadata inton the database");
        }

        // Start uploading the file

        try
        {
            await _fileUploadService.StoreFile(formFile, newFileName);
        } 
        catch (EmptyFileException e)
        {
            return BadRequest(e.Message);
        }
        catch
        {
            return StatusCode(
                StatusCodes.Status500InternalServerError, 
                "Issue uploading the file to the server.");
        }
        
        // Return the status of the upload

        return Ok(new FileUploadStatusResponse
        {
            DataTemplateId = dataTemplateId,
            FileId = newFile.Id,
            Status = FileUploadStatus.Uploading,
        });
    }
}