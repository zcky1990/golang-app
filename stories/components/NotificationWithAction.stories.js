import { fn } from "@storybook/test";
import notificationWithAction from "../../static/javascript/components/NotificationWithAction.vue";

export default {
  title: "Components/NotificationWithAction",
  component: notificationWithAction,
  parameters: {
    docs: {
      description: {
        component: 'A Notification dialog component that show message with button selection. Customize the dialog using props.'
      }
    }
  },
  tags: ["autodocs"],
  render: (args) => ({
    components: {
      notificationWithAction,
    },
    setup() {
      return {
        ...args,
      };
    },
    template:
      '<notificationWithAction :showDialog="showDialog" :titleDialog="titleDialog" :messageDialog="messageDialog"  :typeDialog="typeDialog" :buttonDialog="buttonDialog" @submitCallback="submitCallback" @cancelCallback="cancelCallback" />',
  }),
  argTypes: {
    showDialog: { control: "boolean" },
    titleDialog: { control: "text" },
    messageDialog: { control: "text" },
    buttonDialog: { control: "object" },
    typeDialog: { control: "text" },
  },
  args: {
    submitCallback: fn(),
    cancelCallback: fn(),
  },
};

export const Notification = {
  parameters: {
    docs: {
      description: {
        story: 'Example how notification dialog will show when type is notification'
      },
      source: {
        code: `<notificationWithAction :showDialog="true" titleDialog="Notification" messageDialog="Notification, Lorem ipsum dolor sit amet consectetur adipisicing elit."  
        typeDialog="notification" :buttonDialog="buttonDialog" @submitCallback="submitCallback" @cancelCallback="cancelCallback" />`
      },
    }
  },
  args: {
    showDialog: true,
    titleDialog: "Notification",
    messageDialog: "Notification, Lorem ipsum dolor sit amet consectetur adipisicing elit.",
    buttonDialog: {
      submit: "Ok",
      cancel: "Cancel",
    },
    typeDialog: "notification",
  },
};

export const Warning = {
  parameters: {
    docs: {
      description: {
        story: 'Example how notification dialog will show when type is warning'
      },
      source: {
        code: `<notificationWithAction :showDialog="true" titleDialog="Warning" messageDialog="Warning, Lorem ipsum dolor sit amet consectetur adipisicing elit."  
        typeDialog="warning" :buttonDialog="buttonDialog" @submitCallback="submitCallback" @cancelCallback="cancelCallback" />`
      },
    }
  },
  args: {
    showDialog: true,
    titleDialog: "Warning",
    messageDialog: "Warning, Lorem ipsum dolor sit amet consectetur adipisicing elit.",
    buttonDialog: {
      submit: "Ok",
      cancel: "Cancel",
    },
    typeDialog: "warning",
  },
};
