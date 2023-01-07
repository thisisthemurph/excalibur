namespace Excalibur.Api.Controllers;

using AutoMapper;
using Excalibur.Api.DTOs.Requests;
using Excalibur.Api.DTOs.Responses;
using Excalibur.Api.Models;
using Excalibur.Api.Services;
using Microsoft.AspNetCore.Mvc;
using SharpCompress.Common;

[ApiController]
[Route("[controller]")]
public class DataTemplateController : ControllerBase
{
    private readonly IMapper _mapper;
    private readonly ILogger<DataTemplateController> _logger;
    private readonly DataTemplateService _dataTemplateService;

    public DataTemplateController(
        IMapper mapper,
        ILogger<DataTemplateController> logger,
        DataTemplateService dataTemplateService)
    {
        _mapper = mapper;
        _logger = logger;
        _dataTemplateService = dataTemplateService;
    }

    [HttpGet]
    public async Task<IEnumerable<DataTemplateResponse>> Get()
    {
        var templates = await _dataTemplateService.GetAsync();
        return templates.Select(t => _mapper.Map<DataTemplateResponse>(t)).ToList();
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<DataTemplateResponse>> GetDataTemplateById(string id, CancellationToken cancellationToken)
    {
        var template = await _dataTemplateService.GetByIdAsync(id, cancellationToken);
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
            var created = await _dataTemplateService.CreateAsync(_mapper.Map<DataTemplate>(dataTemplate));
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

        var success = await _dataTemplateService.AddColumnAsync(id, column, cancellationToken);
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
            var updatedDataTemplate = await _dataTemplateService.UpdateAsync(id, dataTemplate.Name, cancellationToken)
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
        var success = await _dataTemplateService.DeleteAsync(id, cancellationToken);
        if (!success)
        {
            return NotFound($"DataTemplate with ID `{id}` does not exist");
        }

        return NoContent();
    }
}