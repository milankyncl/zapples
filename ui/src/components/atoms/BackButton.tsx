import React, {FC} from "react";
import {ArrowLeftIcon} from "@heroicons/react/24/solid";

interface Props {
    onClick: () => void;
}

export const BackButton: FC<Props> = ({ onClick }) => {
    return (
        <button onClick={onClick} type="button" className="text-gray-400 border border-gray-400 hover:bg-gray-100 font-medium rounded-lg text-sm p-2 text-center inline-flex items-center mr-5">
            <ArrowLeftIcon className="w-5 h-5" />
            <span className="sr-only">Go back</span>
        </button>
    )
}