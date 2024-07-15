import { fn } from "@storybook/test";
import snackbar from "../../static/javascript/components/Snackbar.vue";

export default {
  title: "Components/SnackBar",
  component: snackbar,
  tags: ["autodocs"],
  parameters: {
    docs: {
      description: {
        component:
          "A Snackbar component that show message. Customize snackbar type by using type props.",
      },
    },
  },
  render: (args) => ({
    components: {
      snackbar,
    },
    setup() {
      return {
        ...args,
      };
    },
    template: `
    <div class="container" style="height:200px;">
<snackbar :show="show" :message="snackbarMessage" :timeout="1000" :type="type" :title="title"  @close="closeSnackbar" />
    </div>
    `,
  }),
  argTypes: {
    snackbarMessage:
      "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?",
    type: {
      options: ["info", "error"],
      control: { type: "select" },
    },
    title: { control: "text" },
    show: {
      options: ["true", "false"],
      control: { type: "select" },
    },
    timeout: {
      control: "number",
    },
    showSnackbar: {
      options: ["true", "false"],
      control: { type: "select" },
    },
    closeSnackbar: {
      options: ["true", "false"],
      control: { type: "select" },
    },
  },
  args: {
    showSnackbar: fn(),
    closeSnackbar: fn(),
  },
};

export const Info = {
  parameters: {
    docs: {
      description: {
        story: "Example how snackbar component renderer",
      },
      source: {
        code: `
        <snackbar 
          :show="true" 
          message="Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?" 
          :timeout=100 
          type="info" 
          title="Info" 
          @close="closeSnackbar" />`,
      },
    },
  },
  args: {
    snackbarMessage:
      "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?",
    show: true,
    type: "info",
    title: "Info",
    timeout: 100,
    showSnackbar: "false",
    closeSnackbar: "false",
  },
};

export const Error = {
  parameters: {
    docs: {
      description: {
        story: "Example how snackbar component renderer",
      },
      source: {
        code: `
        <snackbar 
          :show="true" 
          message="Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?" 
          :timeout=100 
          type="error" 
          title="Error" 
          @close="closeSnackbar" />`,
      },
    },
  },
  args: {
    snackbarMessage:
      "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?",
    show: true,
    type: "error",
    title: "Error",
    timeout: 100,
  },
};
