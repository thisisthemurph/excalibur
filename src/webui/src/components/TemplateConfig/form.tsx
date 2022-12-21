import { useNavigate } from "react-router-dom";
import { useForm, useFieldArray, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "react-query";

import { XMarkIcon } from "@heroicons/react/24/solid";
import { PlusCircleIcon } from "@heroicons/react/24/solid";

import { ColumnDataTypeEnum, DataTemplate, DataTemplateSchema } from "../../types";
import { HateoasResponse } from "../../api/types";
import { deleteDataTemplate } from "../../api/dataTemplate";

type DataTemplateFormSubmitFn = (dt: DataTemplate) => Promise<HateoasResponse>;

interface Props {
	config: DataTemplate;
	onSubmitFn: DataTemplateFormSubmitFn;
	controls?: boolean;
}

const TemplateConfigForm = ({ config, controls: controls, onSubmitFn }: Props) => {
	const navigate = useNavigate();
	const updateMutation = useMutation(onSubmitFn);
	const deleteMutation = useMutation(deleteDataTemplate);

	const hasControls = controls ? true : false;

	const {
		control,
		register,
		handleSubmit,
		formState: { errors, isSubmitting },
	} = useForm<DataTemplate>({
		resolver: zodResolver(DataTemplateSchema),
		defaultValues: config,
	});

	const { fields, append, remove } = useFieldArray({
		control,
		name: "columns",
		rules: { minLength: 1 },
	});

	const onSubmit: SubmitHandler<DataTemplate> = async (data) => {
		await updateMutation.mutateAsync({ ...data });
	};

	const onDelete = async (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
		e.preventDefault();

		if (!config._id) {
			return;
		}

		if (confirm("Are you sure you would like to delete this data template")) {
			await deleteMutation.mutateAsync(config._id);
			navigate("/template");
		}
	};

	const onAppendColumn = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
		e.preventDefault();
		append({
			dataType: ColumnDataTypeEnum.Values.String,
			originalName: "",
			prettyName: "",
		});
	};

	return (
		<form className="my-16 space-y-6 px-wrap" onSubmit={handleSubmit(onSubmit)}>
			<h3>Data table configuration</h3>

			<pre>{JSON.stringify({ errors, isSubmitting }, null, 2)}</pre>

			<section>
				<label htmlFor="name" className="inline-block pb-2 text-gray-700">
					Table name
				</label>
				<input
					type="text"
					id="tableName"
					className="w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
					{...register("name")}
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
						<th className="pb-2 pl-2 text-left font-normal text-gray-700">Data type</th>
						<th className="pb-2 pl-2 text-left font-normal text-gray-700">Original name</th>
						<th className="pb-2 pl-2 text-left font-normal text-gray-700">Pretty name</th>
						<th className="pb-2 pl-2 text-left font-normal text-gray-700">&nbsp;</th>
					</tr>
				</thead>

				<tbody>
					{fields.map(({ id, dataType, originalName, prettyName }, index) => (
						<tr key={id} className="my-2">
							<td>
								<select {...register(`columns.${index}.dataType`)} defaultValue={dataType}>
									{ColumnDataTypeEnum.options.map((t) => (
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

			<footer className="flex justify-between gap-1">
				<button className="px-2" onClick={(e) => onAppendColumn(e)}>
					<PlusCircleIcon className="h-12 w-12 text-indigo-400 hover:text-indigo-600" />
				</button>

				<div className="space-x-4">
					{hasControls && (
						<button className="btn btn__danger" onClick={onDelete}>
							{deleteMutation.status === "loading" ? "Deleting" : "Delete"}
						</button>
					)}
					<input
						type="submit"
						className="btn btn__primary"
						value={updateMutation.status === "loading" ? "Saving" : "Save"}
					/>
				</div>
			</footer>
		</form>
	);
};

export default TemplateConfigForm;
