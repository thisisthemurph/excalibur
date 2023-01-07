namespace Excalibur.Api.Controllers;

using Excalibur.Api.DTOs;
using Excalibur.Api.DTOs.Requests;
using Excalibur.Api.DTOs.Responses;
using Excalibur.Api.Services;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("[controller]")]
public class DataTemplateController : ControllerBase
{
    private readonly ILogger<DataTemplateController> _logger;
    private readonly DataTemplateService _mongoDBService;

    public DataTemplateController(
        ILogger<DataTemplateController> logger,
        DataTemplateService mongoDbService)
    {
        _logger = logger;
        _mongoDBService = mongoDbService;
    }

    [HttpGet]
    public async Task<IEnumerable<DataTemplateResponse>> Get()
    {
        var templates = await _mongoDBService.GetAsync();
        return templates.Select(t => new DataTemplateResponse
        {
            Id = t.Id,
            Name= t.Name,
            Columns = t.Columns is null ? new List<DataTemplateColumnDto>() : t.Columns.Select(c => new DataTemplateColumnDto { Id = c.Id.ToString(), DataType = c.DataType, OriginalName = c.OriginalName, PrettyName = c.PrettyName }).ToList(),
            Files = t.Files,
        }).ToList();
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<DataTemplateResponse>> GetDataTemplateById(string id, CancellationToken cancellationToken)
    {
        var template = await _mongoDBService.GetByIdAsync(id, cancellationToken);
        if (template is null)
        {
            return NotFound($"DataTemplate with ID `{id}` does not exist");
        }

        return Ok(template.MapToResponse());
    }

    [HttpPost]
    public async Task<IActionResult> Create([FromBody] DataTemplateCreateRequest dataTemplate)
    {
        if (!ModelState.IsValid)
        {
            return BadRequest(ModelState);
        }

        var created = await _mongoDBService.CreateAsync(dataTemplate);
        return Created(nameof(Get), new { id = created.Id });
    }

    [HttpPost("{id}")]
    public async Task<IActionResult> AddColumn(string id, [FromBody] DataTemplateCreateColumnRequest column, CancellationToken cancellationToken)
    {
        if(!ModelState.IsValid)
        {
            return BadRequest(ModelState);
        }

        var success = await _mongoDBService.AddColumnAsync(id, column, cancellationToken);
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

        var updated = await _mongoDBService.UpdateAsync(id, dataTemplate.Name, cancellationToken);

        if (updated is null)
        {
            return NotFound($"DataTemplate with ID {id} does not exist");
        }

        return Ok(updated.MapToResponse());
    }

    [HttpDelete("{id}")]
    public async Task<IActionResult> Delete(string id, CancellationToken cancellationToken)
    {
        var success = await _mongoDBService.DeleteAsync(id, cancellationToken);
        if (!success)
        {
            return NotFound($"DataTemplate with ID `{id}` does not exist");
        }

        return NoContent();
    }
}