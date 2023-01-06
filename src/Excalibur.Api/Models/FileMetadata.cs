namespace Excalibur.Api.Models;

public class FileMetadata
{
    public string Id { get; set; }
    public string Name { get; set; }
    public string StoredName { get; set; }
    public string Status { get; set; } = FileUploadStatus.Unknown;
}

public static class FileUploadStatus
{
    // Unknown - the status of the file is not known
    public static readonly string Unknown = "Unknown";

    // Uploading - a file that is currently being uploaded
    public static readonly string Uploading = "Uploading";

    // Complete - a file that has successfully uploaded
    public static readonly string Complete = "Complete";

    // Failed - the upload has failed
    public static readonly string Failed = "Failed";
}
