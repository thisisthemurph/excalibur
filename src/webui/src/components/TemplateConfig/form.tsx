import { useForm, useFieldArray, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  DataTypeEnum,
  defaultColumnObject,
  FormColumnSchemaType,
  FormSchema,
  FormSchemaType,
} from "./z";
import Button from "../Button";

interface Props {
  config: FormSchemaType;
}

const TemplateConfigForm = ({ config }: Props) => {
  const {
    control,
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FormSchemaType>({
    resolver: zodResolver(FormSchema),
    defaultValues: config,
  });

  const { fields, append, remove } = useFieldArray({
    control,
    name: "columns",
    rules: { minLength: 1 },
  });

  const onSubmit: SubmitHandler<FormSchemaType> = async (data) => {
    console.log(data);
  };

  const onAppendColumn = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();
    append(defaultColumnObject);
  };

  return (
    <form className="config form" onSubmit={handleSubmit(onSubmit)}>
      <h3>Data table configuration</h3>

      <section className="form__section table-name">
        <label htmlFor="tableName">Table name: </label>
        <input id="tableName" {...register("tableName")} />
      </section>

      <table className="config-table">
        <colgroup>
          <col className="col__dataType" />
          <col className="col__originalName" />
          <col className="col__prettyName" />
          <col className="col__buttons" />
        </colgroup>

        <thead>
          <tr>
            <th>Data type</th>
            <th>Original name</th>
            <th>Pretty name</th>
            <th>&nbsp;</th>
          </tr>
        </thead>

        <tbody>
          {fields.map(({ id, dataType, originalName, prettyName }, index) => (
            <tr key={id}>
              <td>
                <select {...register(`columns.${index}.dataType`)} defaultValue={dataType}>
                  {DataTypeEnum.options.map((t) => (
                    <option key={`${t}-${id}`} value={t}>
                      {t}
                    </option>
                  ))}
                </select>
              </td>
              <td>
                <input {...register(`columns.${index}.originalName`)} defaultValue={originalName} />
              </td>
              <td>
                <input {...register(`columns.${index}.prettyName`)} defaultValue={prettyName} />
              </td>

              <td className="buttons">
                <button className="btn btn__small btn__danger" onClick={() => remove(index)}>
                  -
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      <footer className="form__button-group">
        <button className="btn btn__small" onClick={(e) => onAppendColumn(e)}>
          +
        </button>
        <Button variant="Submit" text="Create"></Button>
      </footer>
    </form>
  );
};

export default TemplateConfigForm;
