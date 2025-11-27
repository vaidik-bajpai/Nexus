import { Field, Input } from "@chakra-ui/react";
import { UseFormRegisterReturn } from "react-hook-form";

interface FormFieldProps {
    label: string;
    placeholder?: string;
    type?: string;
    register: UseFormRegisterReturn;
    error?: { message?: string };
}

const FormField = ({
    label,
    placeholder,
    type = "text",
    register,
    error,
}: FormFieldProps) => {
    return (
        <Field.Root invalid={!!error}>
            <Field.Label>{label}</Field.Label>
            <Input placeholder={placeholder} type={type} {...register} />
            <Field.ErrorText>{error?.message}</Field.ErrorText>
        </Field.Root>
    );
};

export default FormField;
