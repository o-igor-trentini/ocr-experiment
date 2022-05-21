import { ChatTeardropDots } from 'phosphor-react-native';
import React, { FC } from 'react';
import { TouchableOpacity } from 'react-native';
import { theme } from '../../theme';
import { styles } from './styles';

interface GetImagesButtonProps {
    onPress: () => void;
}

export const GetImagesButton: FC<GetImagesButtonProps> = ({ onPress }) => {
    return (
        <>
            <TouchableOpacity onPress={onPress} style={styles.button}>
                <ChatTeardropDots size={24} color={theme.colors.text_on_brand_color} weight="bold" />
            </TouchableOpacity>
        </>
    );
};
