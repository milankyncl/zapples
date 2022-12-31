import React, {FC} from 'react';
import {SubmitHandler, useForm} from "react-hook-form";
import {Button, ButtonSize} from "../atoms/Button";

interface Props {
    loading: boolean;
    onSubmit: (_: FeatureFormDto) => void;
}

export interface FeatureFormDto {
    key: string;
    description: string;
}

export const FeatureForm: FC<Props> = ({
    onSubmit,
    loading,
}) => {
    const { register, handleSubmit, formState: { isValid } } = useForm<FeatureFormDto>();

    const submitHandler: SubmitHandler<FeatureFormDto> = (data) => {
        onSubmit(data);
    };

    return (
        <form onSubmit={handleSubmit(submitHandler)}>
            <div className="mb-4">
                <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="key">
                    Feature key
                </label>
                <input
                    type="text"
                    {...register('key')}
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                />
            </div>
            <div>
                <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="description">
                    Description
                </label>
                <textarea
                    {...register('description')}
                    rows={5}
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                />
            </div>
            <div className="text-center mt-4">
                <Button type="submit" size={ButtonSize.Lg} disabled={!isValid || loading}>
                    Create feature
                </Button>
            </div>
        </form>
    )
}