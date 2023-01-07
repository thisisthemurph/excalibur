using System.Globalization;

namespace Excalibur.Api.ExtensionMethods;

public static class DateTimeExtensionMethods
{
    /// <summary>
    /// Returns a string representation of the DateTime object in the ISO 8601 format.
    /// </summary>
    public static string ToIsoFormatString(this DateTime dateTime)
    {
        return dateTime.ToString("yyyy-MM-ddTHH\\:mm\\:ss", CultureInfo.InvariantCulture);
    }
}
