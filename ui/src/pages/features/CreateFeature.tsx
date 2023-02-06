import React, {FC, useState} from 'react';
import {useNavigate} from "react-router-dom";
import {BackButton} from "../../components/atoms/BackButton";
import {PageHeading} from "../../components/atoms/PageHeading";
import {FeatureForm, FeatureFormDto} from "../../components/features/FeatureForm";
import {client} from "../../api/client";
import moment from "moment";

export const CreateFeature: FC = () => {
    const navigate = useNavigate();
    const [loading, setLoading] = useState(false);

    const onSubmitForm = async (data: FeatureFormDto) => {
        setLoading(true);
        try {
            await client.post('/features', {
                key: data.key,
                description: data.description,
                enabledSince: data.enabledSince ? moment(data.enabledSince).toISOString() : null,
                enabledUntil: data.enabledUntil ? moment(data.enabledUntil).toISOString() : null,
            })
        } catch (e) {
            console.error(e, 'Error happened while creating features')
        } finally {
            setLoading(false);
            navigate('/');
        }
    };

    return <>
        <div className="flex items-center mb-2">
            <BackButton onClick={() => navigate('/')}></BackButton>
            <PageHeading>Create feature</PageHeading>
        </div>
        <FeatureForm loading={loading} onSubmit={onSubmitForm} />
    </>
}