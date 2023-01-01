import React, {FC, useState} from 'react';
import {useNavigate} from "react-router-dom";
import {BackButton} from "../../components/atoms/BackButton";
import {PageHeading} from "../../components/atoms/PageHeading";
import {FeatureForm, FeatureFormDto} from "../../components/features/FeatureForm";
import {client} from "../../api/client";

export const CreateFeature: FC = () => {
    const navigate = useNavigate();
    const [loading, setLoading] = useState(false);

    const onSubmitForm = async (data: FeatureFormDto) => {
        setLoading(true);
        try {
            await client.post('/features', {
                key: data.key,
                description: data.description,
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