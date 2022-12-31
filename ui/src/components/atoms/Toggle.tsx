import React, {FC} from "react";
import classNames from "classnames";

interface Props {
    defaultToggled: boolean;
    onToggle: (toggled: boolean) => void;
    disabled?: boolean;
}

export const Toggle: FC<Props> = ({
    defaultToggled,
    onToggle,
    disabled,
}) => {
    return (
        <label className={classNames("inline-flex relative items-center cursor-pointer transition-opacity", disabled ? 'opacity-50' : '')}>
            <input type="checkbox" className="sr-only peer" defaultChecked={defaultToggled} onChange={(e) => onToggle(e.currentTarget.checked)} disabled={disabled} />
            <div className="w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600" />
        </label>
    )
}