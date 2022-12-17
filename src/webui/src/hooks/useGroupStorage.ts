import { useCallback, useEffect, useState } from "react";
import useLocalStorage from "./useLocalStorage";

type ItemObjectType<T> = { [key: string]: T };

// key: unique identifier of the group of items
// itemId: unique item of a specific item within the group
// defaultItem: the default item should the itemId not exist in the group
function useGroupStorage<T>(key: string, itemId: string, defaultItem: T): [T, (value: T) => void] {
	const [storedValues, setStoredValues] = useLocalStorage<ItemObjectType<T>>(key, {});
	const [currentValue, setCurrentValue] = useState<T>(storedValues[itemId] || defaultItem);

	const setValue = useCallback((value: T) => {
		const updatedValues = { ...storedValues, [itemId]: value };

		setCurrentValue(value);
		setStoredValues(updatedValues);
	}, []);

	useEffect(() => {
		setValue(defaultItem);
	}, []);

	return [currentValue || defaultItem, setValue];
}

export default useGroupStorage;
