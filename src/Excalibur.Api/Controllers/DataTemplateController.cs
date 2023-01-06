namespace Excalibur.Api.Controllers;

using Excalibur.Api.DTOs;
using Excalibur.Api.Models;
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
    public async Task<IEnumerable<DataTemplateDto>> Get()
    {
        var templates = await _mongoDBService.GetAsync();
        return templates.Select(t => t.MapToDto()).ToList();
    }

    [HttpGet("{id}")]
    public async Task<IActionResult> GetDataTemplateById(string id, CancellationToken cancellationToken)
    {
        var template = await _mongoDBService.GetByIdAsync(id, cancellationToken);
        if (template is null)
        {
            return NotFound($"DataTemplate with ID `{id}` does not exist");
        }

        return Ok(template.MapToDto());
    }

    [HttpPost]
    public async Task<IActionResult> Create([FromBody] DataTemplateDto dataTemplate)
    {
        if (!ModelState.IsValid)
        {
            return BadRequest(ModelState);
        }

        var created = await _mongoDBService.CreateAsync(dataTemplate.MapToModel());
        return Created(nameof(Get), new { id = created.Id });
    }

    [HttpPost("{id}")]
    public async Task<IActionResult> AddColumn(string id, [FromBody] DataTemplateColumnCreateDto column, CancellationToken cancellationToken)
    {
        if(!ModelState.IsValid)
        {
            return BadRequest(ModelState);
        }

        var success = await _mongoDBService.AddColumnAsync(id, column.MapToModel(), cancellationToken);
        if (!success)
        {
            return NotFound($"DataTemplate with ID `{id}` does not exist");
        }

        return NoContent();
    }

    [HttpPut("{id}")]
    public async Task<ActionResult<DataTemplateDto>> Update(string id, [FromBody] DataTemplateUpdateDto dataTemplate, CancellationToken cancellationToken)
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

        return Ok(updated.MapToDto());
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