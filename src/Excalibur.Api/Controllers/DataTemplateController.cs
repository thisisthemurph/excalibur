namespace Excalibur.Api.Controllers;

using AutoMapper;
using Excalibur.Application.DTOs.Requests;
using Excalibur.Application.DTOs.Responses;
using Excalibur.Application.Repositories;
using Excalibur.Domain.Entities;
using Excalibur.Domain.Models;
using Microsoft.AspNetCore.Mvc;
using SharpCompress.Common;

[ApiController]
[Route("[controller]")]
public class DataTemplateController : ControllerBase
{
    private readonly IMapper _mapper;
    private readonly ILogger<DataTemplateController> _logger;
    private readonly IDataTemplateRepo _dataTemplateRepo;

    public DataTemplateController(
        IMapper mapper,
        ILogger<DataTemplateController> logger,
        IDataTemplateRepo dataTemplateRepo)
    {
        _mapper = mapper;
        _logger = logger;
        _dataTemplateRepo = dataTemplateRepo;
    }

    [HttpGet]
    public async Task<IEnumerable<DataTemplateResponse>> Get()
    {
        var templates = await _dataTemplateRepo.GetAsync();
        return templates.Select(t => _mapper.Map<DataTemplateResponse>(t)).ToList();
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<DataTemplateResponse>> GetDataTemplateById(string id, CancellationToken cancellationToken)
    {
        var template = await _dataTemplateRepo.GetByIdAsync(id, cancellationToken);
        if (template is null)
        {
            return NotFound($"DataTemplate with ID `{id}` does not exist");
        }

        return Ok(_mapper.Map<DataTemplateResponse>(template));
    }

    [HttpPost]
    public async Task<IActionResult> Create([FromBody] DataTemplateCreateRequest dataTemplate)
    {
        if (!ModelState.IsValid)
        {
            return BadRequest(ModelState);
        }

        try
        {
            var created = await _dataTemplateRepo.CreateAsync(_mapper.Map<DataTemplate>(dataTemplate));
            return Created(nameof(Get), new { id = created.Id });
        }
        catch (ArgumentException e)
        {
            return BadRequest(e.Message);
        }
        catch (Exception e)
        {
            return StatusCode(StatusCodes.Status500InternalServerError, e);
        }
    }

    [HttpPost("{id}")]
    public async Task<IActionResult> AddColumn(string id, [FromBody] DataTemplateCreateColumnRequest column, CancellationToken cancellationToken)
    {
        if(!ModelState.IsValid)
        {
            return BadRequest(ModelState);
        }

        var success = await _dataTemplateRepo.AddColumnAsync(id, column, cancellationToken);
        if (!success)
        {
            return NotFound($"DataTemplate with ID `{id}` does not exist");
        }

        return NoContent();
    }

    [HttpPut("{id}")]
    public async Task<ActionResult<DataTemplateResponse>> Update(string id, [FromBody] DataTemplateUpdateRequest dataTemplate, CancellationToken cancellationToken)
    {
        if (!ModelState.IsValid)
        {
            return BadRequest(ModelState);
        }

        try
        {
            var updatedDataTemplate = await _dataTemplateRepo.UpdateAsync(id, dataTemplate.Name, cancellationToken);
            if (updatedDataTemplate is null)
            {
                return NotFound($"DataTemplate with ID {id} does not exist");
            }
        }
        catch (ArgumentException e)
        {
            return BadRequest(e.Message);
        }
        catch (Exception e)
        {
            return StatusCode(StatusCodes.Status500InternalServerError, e);
        }

        return NoContent();
    }

    [HttpDelete("{id}")]
    public async Task<IActionResult> Delete(string id, CancellationToken cancellationToken)
    {
        var success = await _dataTemplateRepo.DeleteAsync(id, cancellationToken);
        if (!success)
        {
            return NotFound($"DataTemplate with ID `{id}` does not exist");
        }

        return NoContent();
    }

    [HttpGet("{dataTemplateId}/upload/{fileId}/status")]
    public async Task<ActionResult<FileUploadStatusResponse>> GetFileUploadStatus(string dataTemplateId, string fileId, CancellationToken cancellationToken)
    {
        var dataTemplate = await _dataTemplateRepo.GetByIdAsync(dataTemplateId, cancellationToken);
        if (dataTemplate is null)
        {
            _logger.LogWarning("DataTemplate with Id {DataTemplateId} does not exist.", dataTemplateId);
            return NotFound("The requested data template does not exist");
            
        }

        var file = dataTemplate.Files.FirstOrDefault(f => f.Id == fileId);
        if (file is null)
        {
            _logger.LogWarning("DataTemplate with Id {DataTemplateId} does not contain a file with Id {fileId}.", dataTemplateId, fileId);
            return NotFound("The required file metadata does not exist within the data template.");
        }

        var response = new FileUploadStatusResponse 
        { 
            DataTemplateId = dataTemplateId,
            FileId = fileId,
            Status = file.Status,
        };

        return Ok(response);
    }
}