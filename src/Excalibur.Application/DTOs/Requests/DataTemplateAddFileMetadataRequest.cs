using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Excalibur.Application.DTOs.Requests;
public class DataTemplateAddFileMetadataRequest
{
    public string Name { get; set; }
    public string StoredName { get; set; }
    public string Status { get; set; }
}
