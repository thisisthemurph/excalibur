import { useState } from "react";
import TemplateConfigForm from "../../components/TemplateConfig/form";
import { FormSchemaType } from "../../components/TemplateConfig/z";
import "./ColumnConfig.scss";

const formDefaultConfiguration: FormSchemaType = {
  tableName: "Company registered vehicles",
  columns: [
    {
      dataType: "String",
      originalName: "first_name",
      prettyName: "First Name",
    },
    {
      dataType: "String",
      originalName: "last_name",
      prettyName: "Surname",
    },
    {
      dataType: "String",
      originalName: "vreg",
      prettyName: "Registration Number",
    },
    {
      dataType: "Number",
      originalName: "value",
      prettyName: "Value",
    },
    {
      dataType: "Boolean",
      originalName: "Vehicle currently in use",
      prettyName: "In use?",
    },
  ],
};

const CreateTemplatePage = () => {
  return (
    <>
      <h1>Create a new template</h1>
      <TemplateConfigForm config={formDefaultConfiguration} />
    </>
  );
};

export default CreateTemplatePage;
