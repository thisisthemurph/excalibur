import { useForm, useFieldArray, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { XMarkIcon } from "@heroicons/react/24/solid";
import { PlusCircleIcon } from "@heroicons/react/24/solid";
import { SubmitButton } from "../Button";

import { DataTypeEnum, defaultColumnObject, FormSchema, FormSchemaType } from "./z";

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
		<form className="px-wrap my-16 space-y-6" onSubmit={handleSubmit(onSubmit)}>
			<h3>Data table configuration</h3>

			<section className="form__section table-name">
				<label htmlFor="tableName" className="inline-block pb-2 text-gray-700">
					Table name
				</label>
				<input
					type="text"
					id="tableName"
					className="w-full border-gray-300 rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
					{...register("tableName")}
				/>
			</section>

			<table className="w-full">
				<colgroup>
					<col className="" />
					<col className="" />
					<col className="" />
					<col className="min-w-20" />
				</colgroup>

				<thead>
					<tr>
						<th className="text-gray-700 font-normal text-left pb-2 pl-2">Data type</th>
						<th className="text-gray-700 font-normal text-left pb-2 pl-2">Original name</th>
						<th className="text-gray-700 font-normal text-left pb-2 pl-2">Pretty name</th>
						<th className="text-gray-700 font-normal text-left pb-2 pl-2">&nbsp;</th>
					</tr>
				</thead>

				<tbody>
					{fields.map(({ id, dataType, originalName, prettyName }, index) => (
						<tr key={id} className="my-2">
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
								<input
									type="text"
									{...register(`columns.${index}.originalName`)}
									defaultValue={originalName}
								/>
							</td>
							<td>
								<input
									type="text"
									{...register(`columns.${index}.prettyName`)}
									defaultValue={prettyName}
								/>
							</td>

							<td>
								<section className="flex justify-end">
									<button className="simple danger" onClick={() => remove(index)}>
										<XMarkIcon className="h-6 w-6 text-red-500 hover:scale-150" />
									</button>
								</section>
							</td>
						</tr>
					))}
				</tbody>
			</table>

			<footer className="flex gap-1 justify-between">
				<button className="px-2" onClick={(e) => onAppendColumn(e)}>
					<PlusCircleIcon className="h-12 w-12 text-indigo-400 hover:text-indigo-600" />
				</button>

				<SubmitButton text="Create" />
			</footer>
		</form>
	);
};

export default TemplateConfigForm;
