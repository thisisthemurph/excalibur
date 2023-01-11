using System.ComponentModel;

namespace Excalibur.Domain.Enums;

public enum ColumnDataType
{
	[Description("Text")]
	Text,

	[Description("Number")]
	Number,

	[Description("Boolean")]
	Boolean,
}
