import React, {FC, useEffect} from 'react';
import {useForm} from "react-hook-form";
import {Button, ButtonSize} from "../atoms/Button";
import {Feature} from "../../api/models/feature";
import DateTime from 'react-datetime';
import 'react-datetime/css/react-datetime.css';
import moment from "moment";

interface Props {
    initialData?: Feature;
    loading: boolean;
    onSubmit: (_: FeatureFormDto) => void;
}

export interface FeatureFormDto {
    key: string;
    description: string;
    enabledSince?: Date;
    enabledUntil?: Date;
}

function validateFeatureKey(str: string) {
    return !/[`!@#$%^&*()+=\[\]{};':"\\|,.<>\/?~]/.test(str);
}


export const FeatureForm: FC<Props> = ({
    initialData,
    onSubmit,
    loading,
}) => {
    const { register, handleSubmit, formState: { isValid }, setValue, watch } = useForm<FeatureFormDto>();

    useEffect(() => {
        if (initialData) {
            setValue('key', initialData.key);
            setValue('description', initialData.description || '');
            initialData.enabledSince && setValue('enabledSince', moment(initialData.enabledSince).toDate());
            initialData.enabledUntil && setValue('enabledUntil', moment(initialData.enabledUntil).toDate());
        }
    }, [initialData])

    const onChangeDate = (field: keyof FeatureFormDto) => (date: string | moment.Moment) => {
        if (typeof date !== 'string') {
            setValue(field, date.toDate());
        }
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="mb-4">
                <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="key">
                    Feature key
                </label>
                <input
                    type="text"
                    {...register('key', { required: true, validate: validateFeatureKey })}
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                />
            </div>
            <div className="mb-2">
                <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="description">
                    Description
                </label>
                <textarea
                    {...register('description')}
                    rows={5}
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                />
            </div>
            <div className="mb-4 leading-5 text-gray-400 text-xs">
                <p>
                    Below you can define date range within the feature will be enabled.
                    <br />
                    To restrict feature for date range, you also need to enable it in dashboard.
                </p>
            </div>
            <div className="mb-4">
                <label className="block text-gray-700 text-sm font-bold mb-2">
                    Enabled since
                </label>
                <DateTime
                    closeOnSelect={true}
                    onChange={onChangeDate('enabledUntil')}
                    value={watch('enabledUntil')}
                    inputProps={{ placeholder: 'forever', className: "shadow appearance-none border rounded w-full py-2.5 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline text-sm" }}
                />
            </div>
            <div>
                <label className="block text-gray-700 text-sm font-bold mb-2">
                    Enabled until
                </label>
                <DateTime
                    closeOnSelect={true}
                    onChange={onChangeDate('enabledSince')}
                    value={watch('enabledSince')}
                    inputProps={{ placeholder: 'forever', className: "shadow appearance-none border rounded w-full py-2.5 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline text-sm" }}
                />
            </div>
            <div className="text-center mt-6">
                <Button type="submit" size={ButtonSize.Lg} disabled={!isValid || loading}>
                    Save feature
                </Button>
            </div>
        </form>
    )
}
